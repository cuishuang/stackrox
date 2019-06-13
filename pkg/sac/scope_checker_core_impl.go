package sac

import (
	"context"
	"sync/atomic"

	"github.com/stackrox/default-authz-plugin/pkg/payload"
	"github.com/stackrox/rox/pkg/sync"
)

// ScopeCheckerCoreImpl can represent a Verb, a Noun, a Cluster, or a Namespace
// Verbs contain a map of Nouns, Nouns contain a map of Clusters, and Clusters contain a map of Namespaces
// Each of these is a valid scope and so each will contain a TryAllowedResult
type ScopeCheckerCoreImpl struct {
	childrenLock sync.Mutex
	children     map[ScopeKey]ScopeCheckerCore

	requestedLock sync.Mutex
	requested     bool

	state        int32
	currentScope payload.AccessScope
	reqTracker   ScopeRequestTracker
}

func createSubScope(currentScope payload.AccessScope, subScopeKey ScopeKey) payload.AccessScope {
	subAccessScope := currentScope
	switch t := subScopeKey.(type) {
	case AccessModeScopeKey:
		subAccessScope.Verb = t.Verb()
	case ResourceScopeKey:
		subAccessScope.Noun = subScopeKey.String()
	case ClusterScopeKey:
		subAccessScope.Attributes.Cluster.ID = subScopeKey.String()
	case NamespaceScopeKey:
		subAccessScope.Attributes.Namespace = subScopeKey.String()
	}
	return subAccessScope
}

// SubScopeChecker returns a sub scope for this scope, or this scope if this scope has been allowed
func (scc *ScopeCheckerCoreImpl) SubScopeChecker(scopeKey ScopeKey) ScopeCheckerCore {
	if scc.atomicLoadState() == Allow {
		return allowAllScopeCheckerCore
	}

	scc.childrenLock.Lock()
	defer scc.childrenLock.Unlock()
	subScope, ok := scc.children[scopeKey]
	if !ok {
		subScope = NewScopeCheckerCore(createSubScope(scc.currentScope, scopeKey), scc.reqTracker)
		scc.children[scopeKey] = subScope
		return subScope
	}
	return subScope
}

// TryAllowed returns Allow/Deny/Unknown as per the comment on the interface
func (scc *ScopeCheckerCoreImpl) TryAllowed() TryAllowedResult {
	state := scc.atomicLoadState()
	if state != Unknown {
		return state
	}

	scc.requestedLock.Lock()
	defer scc.requestedLock.Unlock()
	if !scc.requested {
		scc.reqTracker.AddRequested(scc)
		scc.requested = true
	}
	return state
}

// PerformChecks performs all pending permission checks as per the comment on the interface
func (scc *ScopeCheckerCoreImpl) PerformChecks(ctx context.Context) error {
	return scc.reqTracker.PerformChecks(ctx)
}

// SetState sets the Allow/Deny/Unknown state of this ScopeCheckerCore, it should only be called by RootScopeCheckerCore
func (scc *ScopeCheckerCoreImpl) SetState(state TryAllowedResult) {
	scc.atomicStoreState(state)
	if state == Allow {
		// if this scope is allowed then all sub-scopes will be allowed
		scc.children = nil
	}
}

// GetAccessScope returns the scope represented by this node as an AccessScope
func (scc *ScopeCheckerCoreImpl) GetAccessScope() payload.AccessScope {
	return scc.currentScope
}

func (scc *ScopeCheckerCoreImpl) atomicLoadState() TryAllowedResult {
	return TryAllowedResult(atomic.LoadInt32(&scc.state))
}

func (scc *ScopeCheckerCoreImpl) atomicStoreState(newState TryAllowedResult) {
	atomic.StoreInt32(&scc.state, int32(newState))
}
