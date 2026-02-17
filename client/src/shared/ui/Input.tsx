const Input = ({ children }: { children: React.ReactNode }) => {
    return (
        <input className="px-6 py-2 rounded-lg font-semibold shadow-lg hover:shadow-xl transition-shadow bg-gradient-to-r from-primary to-secondary text-white">
            {children}
        </input>
    );
};

export default Input;