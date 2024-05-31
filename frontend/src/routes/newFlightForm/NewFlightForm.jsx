// src/components/NewFlightForm.jsx
import { useState } from 'react';
import { TextInput, Button, Paper, Container } from '@mantine/core';
import { DateInput } from '@mantine/dates'
import { useForm } from '@mantine/form';
import { createFlight } from '../../api/createFlight';

export function NewFlightForm() {
    const [loading, setLoading] = useState(false);
    const form = useForm({
        initialValues: {
            to: '',
            from: '',
            id: '',
            date: '',
        },
    });

    const createFlightHandler = async (values = form.values) => {
        setLoading(true);
        try {
            await createFlight(values);
            form.reset();
        } catch (error) {
            console.error('Error creating flight:', error);
        } finally {
            setLoading(false);
        }
    };

    return (
        <Container size="md">
            <Paper withBorder p="md">
                <form onSubmit={form.onSubmit(createFlightHandler)}>
                    <TextInput
                        label="To"
                        placeholder="Enter destination airport code"
                        required
                        disabled={loading}
                        {...form.getInputProps('to')}
                    />
                    <TextInput
                        label="From"
                        placeholder="Enter departure airport code"
                        required
                        disabled={loading}
                        {...form.getInputProps('from')}
                    />
                    <TextInput
                        label="Flight ID"
                        placeholder="Enter flight ID"
                        required
                        disabled={loading}
                        {...form.getInputProps('id')}
                    />
                    <DateInput
                        label="Date"
                        placeholder="Pick date and time"
                        required
                        disabled={loading}
                        {...form.getInputProps('date')}
                    />
                    <Button type="submit" loading={loading} mt="sm" disabled={loading}>
                        Create Flight
                    </Button>
                </form>
            </Paper>
        </Container>
    );
}
