import {notifications} from "@mantine/notifications";


export const logout = async () => {
    try {
        const response = await fetch(`${window.location.protocol}//${window.location.hostname}${window.location.port}/api/auth/v1/logout`, {
            method: 'GET',

        });
        if (!response.ok) {
            throw new Error("Logout failed");
        }

        notifications.show({
            title: 'Logout successful',
            message: 'You have successfully logged out!',
            color: 'green',
        });

        return response.json();
    } catch (error) {
        notifications.show({
            title: 'Logout failed',
            message: 'An error occurred. Please try again later.',
            color: 'red',
        })
        console.log("error during logout: ", error)
        throw error;
    }

}