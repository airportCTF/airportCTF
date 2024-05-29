import React from 'react';
import { AppShell, Group, Image } from '@mantine/core';
import logo from '../../assets/AirportLogo.svg';
import { useNavigate } from 'react-router-dom';
import {HeroImageRight} from "./heroWithImageOnTheRight.jsx";
import {ActionToggle} from "../../ActionToggle.jsx";



const Home = () => {
    const navigate = useNavigate();


    return (
        <AppShell header={{height: 60}} >
            <AppShell.Header justify="left" p="xs">
                <Group justify="space-between" >
                    <Image w={130} src={logo} alt="AirportLogo" />
                    <ActionToggle/>
                </Group>

            </AppShell.Header>
            <AppShell.Main>
                <HeroImageRight/>
            </AppShell.Main>
        </AppShell>

    );
};

export default Home;
