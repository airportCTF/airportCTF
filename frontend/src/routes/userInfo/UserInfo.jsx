// src/routes/profile/UserInfo.jsx
import React, { useState, useEffect } from 'react';
import { Avatar, Text, Button, Paper, Skeleton } from '@mantine/core';
import { getProfile } from '../../api/getProfile';

export function UserInfo() {
    const [profile, setProfile] = useState(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchProfile = async () => {
            try {
                const profileData = await getProfile();
                setProfile(profileData);
            } catch (error) {
                console.error('Failed to fetch profile:', error);
            } finally {
                setLoading(false);
            }
        };

        fetchProfile();
    }, []);

    if (loading) {
        return (
            <Skeleton height={200} circle mb="xl">
                <Skeleton height={8} radius="xl" />
                <Skeleton height={8} mt={6} radius="xl" width="70%" />
            </Skeleton>
        );
    }

    if (!profile) {
        return null; // или можно вернуть сообщение об ошибке
    }

    return (
        <Paper radius="md" withBorder p="lg" bg="var(--mantine-color-body)">
            <Avatar
                src="https://raw.githubusercontent.com/mantinedev/mantine/master/.demo/avatars/avatar-8.png"
                size={120}
                radius={120}
                mx="auto"
            />
            <Text ta="center" fz="lg" fw={500} mt="md">
                {profile.name} {profile.lastName}
            </Text>
            <Text ta="center" c="dimmed" fz="sm">
                {profile.login}
            </Text>
            <Text ta="center" c="dimmed" fz="sm">
                Passport Number: {profile.passportNum}
            </Text>
            <Text ta="center" c="dimmed" fz="sm">
                {(profile.isAdmin === "true") ? "Administrator" : "Regular user"}
            </Text>
        </Paper>
    );
}
