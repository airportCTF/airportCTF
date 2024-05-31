// src/api/createFlight.js
import { notifications } from '@mantine/notifications';

export const createFlight = async (flightData) => {
    try {
        const response = await fetch(`${window.location.protocol}//${window.location.hostname}${window.location.port}/api/controlroom/v1/admin/flight/new`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(flightData),
        });

        if (!response.ok) {
            throw new Error(`Create flight ${flightData.id} failed`);
        }

        notifications.show({
            title: `Flight ${flightData.id} created successfully`,
            message: `Flight from ${flightData.from} to ${flightData.to} on ${flightData.date} is created`,
            color: 'green',
        });

        return response.json();
    } catch (error) {
        notifications.show({
            title: `Error creating flight ${flightData.id}`,
            message: error.message,
            color: 'red',
        });
        throw error;
    }
};
