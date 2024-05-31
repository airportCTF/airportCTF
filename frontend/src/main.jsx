import React from 'react'
import ReactDOM from 'react-dom/client'
import {createBrowserRouter, Navigate, RouterProvider} from "react-router-dom";
import '@mantine/core/styles.css';
import '@mantine/notifications/styles.css';
import '@mantine/dates/styles.css';
import {MantineProvider} from "@mantine/core";
import {Notifications} from "@mantine/notifications";
import Profile from './routes/profile/profile.jsx'
import Home from "./routes/home/home.jsx";
import Tickets from "./routes/tickets/tickets.jsx";
import AuthenticationImage from "./routes/login/login.jsx";
import {NotFoundTitle} from "./Error404Page.jsx";
import {UserInfo} from "./routes/userInfo/UserInfo.jsx";
import {MakeAdminForm} from "./routes/makeAdminForm/makeAdminForm.jsx";
import {NewFlightForm} from "./routes/newFlightForm/NewFlightForm.jsx";


const router = createBrowserRouter([
    {
        path: "/",
        element: <Home/>,
        errorElement: <NotFoundTitle/>
    },
    {
        path: "profile",
        element: <Profile/>,
        children: [
            { path: "tickets", element: <Tickets/> },
            { path: "userinfo", element: <UserInfo /> },
            { path: "makeAdmin", element: <MakeAdminForm /> },
            { path: "newFlight", element: <NewFlightForm /> },
        ],
        errorElement: <NotFoundTitle/>
    },
    // {
    //     path: "tickets",
    //     element: <Tickets />,
    // },
    {
        path: "auth",
        element: <AuthenticationImage/>,
        errorElement: <NotFoundTitle/>,
    }

]);


ReactDOM.createRoot(document.getElementById('root')).render(
    <React.StrictMode>
        <MantineProvider defaultColorScheme="light">
            <Notifications zIndex={1000} limit={5} autoClose={4000} />
            <RouterProvider router={router}/>
        </MantineProvider>
    </React.StrictMode>,
)
