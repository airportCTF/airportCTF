import { Container, Title, Text, Button } from '@mantine/core';
import { Link } from 'react-router-dom';
import classes from './HeroImageRight.module.css';

export function HeroImageRight() {
    return (
        <div className={classes.root}>
            <Container size="lg">
                <div className={classes.inner}>
                    <div className={classes.content}>
                        <Title className={classes.title}>
                            <Text
                                component="span"
                                inherit
                                variant="gradient"
                                gradient={{ from: '#40c9ff', to: '#e81cff' }}
                            >
                                Airport
                            </Text>{' '}
                            Service
                        </Title>

                        <Text className={classes.description} mt={30}>
                            It is an airport service for Attack / Defense competitions, providing such features as login and registration, a separate flight tracking board, a personal account, ticket purchase and a control room.
                        </Text>
                        <Link to={"/auth"}>
                            <Button
                                variant="gradient"
                                gradient={{ from: '#40c9ff', to: '#e81cff' }}
                                size="xl"
                                className={classes.control}
                                mt={40}
                            >
                                Get started
                            </Button>
                        </Link>


                    </div>
                </div>
            </Container>
        </div>
    );
}