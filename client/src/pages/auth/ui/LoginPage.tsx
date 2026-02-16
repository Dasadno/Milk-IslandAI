import { LoginForm } from '@/features/auth/ui/LoginForm/LoginForm';

export const LoginPage = () => {
    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-50 dark:bg-gray-900 px-4">
            <LoginForm />
        </div>
    );
};
