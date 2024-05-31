import {notifications} from '@mantine/notifications';

export const getProfile = async () => {
    try {
        const response = await fetch(`${window.location.protocol}//${window.location.hostname}${window.location.port}/api/auth/v1/profile`, {
            method: 'GET',
            credentials: 'include', // для отправки куков вместе с запросом
        });

        if (!response.ok) {
            notifications.show({
                title: 'Error fetching profile',
                message: response.status,
                color: 'red',
            });
            throw new Error('Failed to fetch profile');
        }

        const data = await response.json();

        if (data.status === 'not authorized') {
            notifications.show({
                title: 'Authorization Error',
                message: 'You are not authorized to view this profile.',
                color: 'red',
            });
            throw new Error('Not authorized');
        }

        return data;
    } catch (error) {
        notifications.show({
            title: 'Error fetching profile',
            message: error.message,
            color: 'red',
        });
        throw error;
    }
};
