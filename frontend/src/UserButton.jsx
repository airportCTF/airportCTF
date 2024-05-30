// src/components/UserButton.jsx
import React, { useState, useEffect } from 'react';
import { UnstyledButton, Group, Avatar, Text, Skeleton } from '@mantine/core';
import { IconChevronRight } from '@tabler/icons-react';
import { useNavigate } from 'react-router-dom';
import classes from './UserButton.module.css';
import { getProfile } from './api/getProfile';

export function UserButton() {
    const [profile, setProfile] = useState(null);
    const [loading, setLoading] = useState(true);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchProfile = async () => {
            try {
                const profileData = await getProfile();
                setProfile(profileData);
                console.log(profileData.login);
                localStorage.setItem('user', profileData.login);
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
            <Skeleton height={50} circle mb="xl">
                <Skeleton height={8} radius="xl" />
                <Skeleton height={8} mt={6} radius="xl" width="70%" />
            </Skeleton>
        );
    }

    if (!profile) {
        return null; // или можно вернуть сообщение об ошибке
    }

    return (
        <UnstyledButton className={classes.user} onClick={() => navigate('/profile/userinfo')}>
            <Group>
                <Avatar
                    src="https://raw.githubusercontent.com/mantinedev/mantine/master/.demo/avatars/avatar-8.png"
                    radius="xl"
                />
                <div style={{ flex: 1 }}>
                    <Text size="sm" fw={500}>
                        {profile.name} {profile.lastName}
                    </Text>
                    <Text color="dimmed" size="xs">
                        {profile.login}
                    </Text>
                </div>
                <IconChevronRight style={{ width: '14px', height: '14px' }} stroke={1.5} />
            </Group>
        </UnstyledButton>
    );
}
