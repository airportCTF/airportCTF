import {notifications} from "@mantine/notifications";

export const makeAdmin = async (user, api_token) => {
    try {
        const response = await fetch(`api/auth/v1/make_admin?user=${user}`, {
                method: 'POST',
                headers: {
                    Authorization: `Bearer ${api_token}`,
                }
            }
        );
        if (!response.ok) {
            throw new Error(`Make_admin ${user} failed`);
        }

        notifications.show({
            title: `Make_admin ${user} successful`,
            message: `${user} is admin now`,
            color: 'green',
        })
    } catch (error) {
        notifications.show({
            title: `Error make admin ${user} profile`,
            message: error.message,
            color: 'red',
        })
        throw error;
    }
}