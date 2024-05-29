import React, {useState} from 'react';
import {Anchor, Button, Container, Paper, PasswordInput, Text, TextInput, Title,} from '@mantine/core';
import {upperFirst, useToggle} from '@mantine/hooks';
import {notifications} from '@mantine/notifications';
import {useNavigate} from 'react-router-dom';
import classes from './AuthenticationImage.module.css';

const AuthenticationImage = () => {
        const [type, toggle] = useToggle(['login', 'register']);
        const [login, setLogin] = useState('');
        const [password, setPassword] = useState('');
        const navigate = useNavigate();

        const handleSubmit = async () => {
            const url = type === 'login' ? '/api/auth/v1/login' : '/api/auth/v1/register';
            const requestBody = {login, password};

            try {
                const response = await fetch(url, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(requestBody),
                });

                const data = await response.json();

                if (type === 'login') {
                    if (response.ok) {
                        notifications.show({
                            title: 'Login successful',
                            message: 'You have successfully logged in!',
                            color: 'green',
                        });
                        navigate('/profile');
                    } else {
                        notifications.show({
                            title: 'Login failed',
                            message: 'Invalid login or password.',
                            color: 'red',
                        });
                    }
                }
                if (type === 'register') {
                    console.log(data)
                    if (response.ok) {
                        notifications.show({
                            title: 'Register successful',
                            message: `You have successfully registered profile: ${data}!`,
                            color: 'green',
                        });
                        toggle();
                    } else {
                        notifications.show({
                            title: 'Registration error',
                            message: 'Some error occured! ' + response.status,
                            color: 'red',
                        });
                    }
                }
            } catch
                (error) {
                notifications.show({
                    title: 'Error',
                    message: 'An error occurred. Please try again later.',
                    color: 'red',
                });
            }
        };

        return (
            <Container fluid className={classes.wrapper} m={0} p={0}>
                <Paper className={classes.form} radius={0} p={30}>
                    <Title order={2} className={classes.title} ta="center" mt="md" mb={50}>
                        Welcome back to Airport!
                    </Title>

                    <TextInput
                        label="Login"
                        placeholder="Enter your username"
                        size="md"
                        required
                        value={login}
                        onChange={(event) => setLogin(event.target.value)}
                    />
                    <PasswordInput
                        label="Password"
                        placeholder="Enter your password"
                        mt="md"
                        size="md"
                        required
                        value={password}
                        onChange={(event) => setPassword(event.target.value)}
                    />
                    <Button fullWidth mt="xl" size="md" onClick={handleSubmit}>
                        {upperFirst(type)}
                    </Button>

                    <Text ta="center" mt="md">
                        {type === 'register' ? 'Already have an account?' : "Don't have an account?"}{' '}
                        <Anchor component="button" href="#" fw={700} onClick={() => toggle()} size="xs">
                            {type === 'register' ? 'Login' : 'Register'}
                        </Anchor>
                    </Text>
                </Paper>
            </Container>
        );
    }
;

export default AuthenticationImage;
