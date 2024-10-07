import React from 'react';
import Button from 'react-bootstrap/Button';

const CreateAccountButton = ( { onClick }) => {

    return (
        <Button variant="primary" onClick={onClick}>Create Account</Button>
    )
}

export default CreateAccountButton
