import React, { ReactElement } from 'react';
import { Button } from '@patternfly/react-core';

export type FormResponseMessage = {
    message: string;
    isError: boolean;
} | null;

export type FormSaveButtonProps = {
    children: ReactElement | ReactElement[] | string;
    onSave: () => void;
    isSubmitting: boolean;
    isTesting: boolean;
};

function FormSaveButton({
    children,
    onSave,
    isSubmitting,
    isTesting,
}: FormSaveButtonProps): ReactElement {
    return (
        <Button
            variant="primary"
            onClick={onSave}
            data-testid="create-btn"
            isDisabled={isSubmitting}
            isLoading={isSubmitting && !isTesting}
        >
            {children}
        </Button>
    );
}

export default FormSaveButton;
