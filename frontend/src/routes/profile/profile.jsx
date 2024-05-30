import classes from './navbar.module.css';
import React, {useState} from 'react';
import {AppShell, Group, Image} from '@mantine/core';
import { IconLogout, IconPlane, IconTicket, IconForms } from '@tabler/icons-react';
import {useDisclosure} from "@mantine/hooks";
import {Link, Outlet, useNavigate} from "react-router-dom";
import logo from "../../assets/AirportLogo.svg";
import {ActionToggle} from "../../ActionToggle.jsx";
import {logout} from "../../api/logout.jsx";
import {UserButton} from "../../UserButton.jsx";


const data = [
    {link: '/', label: 'Home page', icon: IconPlane},
    {link: '/profile/tickets', label: 'Tickets', icon: IconTicket},
    {link: '/profile/makeAdmin', label: 'Make Admin', icon: IconForms},
];

export default function Profile() {
    const [active, setActive] = useState('tickets');
    const [opened, {toggle}] = useDisclosure();
    const navigate = useNavigate();

    const links = data.map((item) => (
        <Link
            className={classes.link}
            data-active={item.label === active || undefined}
            to={item.link}
            key={item.label}
            onClick={() => setActive(item.label)}
        >
            <item.icon className={classes.linkIcon} stroke={1.5}/>
            <span>{item.label}</span>
        </Link>
    ));

    return (
        <>
            <AppShell
                header={{height: 60}}
                navbar={{width: 300, breakpoint: 'sm', collapsed: {mobile: !opened}}}
                padding="md">
                <AppShell.Header>
                    <Group className={classes.header} justify="flex-start" p="xs">
                        <Image w={130} src={logo} alt="AirportLogo"/>
                        <ActionToggle/>
                    </Group>
                </AppShell.Header>
                <AppShell.Navbar>
                    <nav className={classes.navbar}>
                        <div className={classes.navbarMain}>
                            {links}
                        </div>
                    </nav>
                    <div className={classes.footer}>
                        <UserButton/>
                    </div>
                        <div className={classes.footer}>
                            <a href="#" className={classes.link} onClick={() => {
                                localStorage.removeItem('user');
                                logout().then(r => console.log(r));
                                navigate('/auth');

                            }}>
                                <IconLogout className={classes.linkIcon} stroke={1.5}/>
                                <span>Logout</span>
                            </a>
                        </div>
                </AppShell.Navbar>
                <AppShell.Main>
                    <Outlet/>
                </AppShell.Main>
            </AppShell>
        </>
    );
}