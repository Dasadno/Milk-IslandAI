import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { HomePage } from '@/pages/home/ui/Page';
import { LoginPage } from '@/pages/auth/ui/LoginPage';
import { ChatPage } from '@/pages/chat/ui/Page';

const router = createBrowserRouter([
    {
        path: '/',
        element: <HomePage />,
    },
    {
        path: '/login',
        element: <LoginPage />,
    },
    {
        path: '/chat',
        element: <ChatPage />,
    },
]);

export const AppRouter = () => {
    return <RouterProvider router={router} />;
};
