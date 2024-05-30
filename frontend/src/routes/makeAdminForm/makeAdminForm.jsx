// src/components/MakeAdminForm.jsx
import { useState } from 'react';
import {TextInput, Button, Paper, Container} from '@mantine/core';
import { useForm } from '@mantine/form';
import { makeAdmin } from '../../api/makeAdmin.jsx';

export function MakeAdminForm() {
    const [loading, setLoading] = useState(false);
    const user = localStorage.getItem('user');
    const form = useForm({
        initialValues: {
            api_token: "",
        }
    });


    const makeAdminHandler = async (values = form.values) => {
        setLoading(true);
        try {
            await makeAdmin(user, values.api_token);
            form.reset();
        } catch (error) {
            console.error('Error making admin:', error);
        } finally {
            setLoading(false);
        }
    };

    return (
        <Container size="md">
            <Paper withBorder p="md" >
                <form onSubmit={form.onSubmit(makeAdminHandler)}>
                    <TextInput
                        key={form.key('api_token')}
                        label="API Token"
                        placeholder="Enter API Token"
                        required
                        disabled={loading}
                        {...form.getInputProps('api_token')}

                    />
                    <Button type="submit" loading={loading} mt="sm" disabled={loading}>
                        Make Admin
                    </Button>
                </form>
            </Paper>
        </Container>


    );
}
