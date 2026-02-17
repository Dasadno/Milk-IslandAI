// import React, { useState } from 'react';

// export const LoginForm = () => {
//     const [email, setEmail] = useState('');
//     const [password, setPassword] = useState('');

//     const handleSubmit = (e: React.FormEvent) => {
//         e.preventDefault();
//         console.log('Login attempt:', { email, password });
//     };

//     return (
//         <form onSubmit={handleSubmit} className="flex flex-col gap-4 w-full max-w-sm p-6 bg-dark-ocean rounded-lg shadow-xl">
//             <h2 className="text-2xl font-bold text-center text-text-primary">Login</h2>

//             <div className="flex flex-col gap-1">
//                 <label htmlFor="email" className="text-sm font-medium text-text-secondary">Email</label>
//                 <input
//                     id="email"
//                     type="email"
//                     value={email}
//                     onChange={(e) => setEmail(e.target.value)}
//                     className="px-4 py-2 bg-deep-midnight border border-dark-ocean rounded-md focus:outline-none focus:ring-2 focus:ring-bright-turquoise text-text-primary"
//                     placeholder="your@email.com"
//                     required
//                 />
//             </div>

//             <div className="flex flex-col gap-1">
//                 <label htmlFor="password" className="text-sm font-medium text-text-secondary">Password</label>
//                 <input
//                     id="password"
//                     type="password"
//                     value={password}
//                     onChange={(e) => setPassword(e.target.value)}
//                     className="px-4 py-2 bg-deep-midnight border border-dark-ocean rounded-md focus:outline-none focus:ring-2 focus:ring-bright-turquoise text-text-primary"
//                     placeholder="********"
//                     required
//                 />
//             </div>

//             <button
//                 type="submit"
//                 className="w-full px-4 py-2 bg-gradient-primary text-white rounded-md hover:shadow-lg focus:outline-none focus:ring-2 focus:ring-bright-turquoise focus:ring-offset-2 focus:ring-offset-deep-midnight transition-shadow font-semibold"
//             >
//                 Sign In
//             </button>
//         </form>
//     );
// };


import React, { useState } from 'react';
import Input from '@/shared/ui/Input';
import { Button } from '@/shared/ui/Button';

export const LoginForm = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        console.log('Login attempt:', { email, password });
    };

    return (
        <form onSubmit={handleSubmit} className="flex flex-col gap-6 w-full">
            {/* Поле Email */}
            <div className="flex flex-col">
                <label className="text-xl font-bold text-white tracking-wide mb-3 ml-1">
                    Email Address
                </label>
                <Input
                    type="email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    placeholder="your@email.com"
                    required
                />
            </div>

            {/* Поле Password */}
            <div className="flex flex-col">
                <label className="text-xl font-bold text-white tracking-wide mb-3 ml-1">
                    Password
                </label>
                <Input
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    placeholder="********"
                    required
                />
            </div>

            <div className="pt-4">
                <Button 
                    variant="gradient" 
                    type="submit" 
                    className="w-full py-4 text-lg uppercase tracking-widest shadow-[0_0_20px_rgba(38,208,206,0.3)]"
                >
                    Войти
                </Button>
            </div>
        </form>
    );
};