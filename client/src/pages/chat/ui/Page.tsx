/**
 * ChatPage - Главная страница чата с AI-агентом
 * 
 * Структура:
 * - Sidebar слева (скрывается на мобилке)
 * - Основная область чата справа
 *   - Header с информацией об агенте
 *   - MessageList (история сообщений)
 *   - MessageInput (поле ввода)
 * 
 * Адаптивность:
 * - Desktop (md+): Sidebar всегда видим, занимает 320px
 * - Mobile (<md): Sidebar скрыт, открывается через кнопку-гамбургер
 * 
 * TODO для декомпозиции:
 * - Вынести ChatHeader в отдельный компонент
 * - Добавить состояние для управления sidebar на мобилке
 * - Добавить анимации открытия/закрытия sidebar
 * - Добавить backdrop для закрытия sidebar при клике вне его
 */

import { useState } from 'react';
import { Menu, X, User, Circle } from 'lucide-react';
import { ChatSidebar, MessageList, MessageInput } from '@/widgets';

export const ChatPage = () => {
    // Состояние для управления видимостью sidebar на мобилке
    const [isSidebarOpen, setIsSidebarOpen] = useState(false);

    // Данные текущего активного агента (потом из API)
    const currentAgent = {
        id: '1',
        name: 'Alice',
        status: 'online',
        personality: 'Curious Explorer',
    };

    return (
        <div className="
            /* Основной контейнер страницы */
            flex
            h-screen
            overflow-hidden
            bg-deep-midnight
        ">
            {/* ========================================
                SIDEBAR (Список агентов)
                ======================================== */}

            {/* Desktop: всегда видим */}
            <div className="hidden md:block">
                <ChatSidebar />
            </div>

            {/* Mobile: показывается через overlay */}
            {isSidebarOpen && (
                <>
                    {/* Backdrop (затемнение фона) */}
                    <div
                        className="
                            fixed inset-0
                            bg-black/50
                            z-40
                            md:hidden
                        "
                        onClick={() => setIsSidebarOpen(false)}
                    />

                    {/* Sidebar с анимацией */}
                    <div className="
                        fixed
                        top-0 left-0
                        h-full
                        w-80
                        z-50
                        md:hidden
                        transform transition-transform duration-300
                    ">
                        <ChatSidebar />
                    </div>
                </>
            )}

            {/* ========================================
                MAIN CHAT AREA (Область чата)
                ======================================== */}
            <div className="
                /* Занимает оставшееся пространство */
                flex-1
                flex flex-col
                overflow-hidden
            ">
                {/* ========================================
                    CHAT HEADER (Шапка чата)
                    ======================================== */}
                <header className="
                    /* Стили шапки */
                    bg-dark-ocean
                    border-b border-bright-turquoise/20
                    px-4 py-3
                    flex items-center gap-4
                    /* Sticky для фиксации при скролле */
                    sticky top-0
                    z-30
                ">
                    {/* Кнопка открытия sidebar (только на мобилке) */}
                    <button
                        onClick={() => setIsSidebarOpen(!isSidebarOpen)}
                        className="
                            md:hidden
                            p-2
                            text-text-primary
                            hover:bg-deep-midnight/50
                            rounded-lg
                            transition-colors
                        "
                    >
                        {isSidebarOpen ? (
                            <X className="w-6 h-6" />
                        ) : (
                            <Menu className="w-6 h-6" />
                        )}
                    </button>

                    {/* Аватар агента (placeholder) */}
                    <div className="
                        w-10 h-10
                        bg-gradient-primary
                        rounded-full
                        flex items-center justify-center
                        text-white
                        font-bold
                    ">
                        {currentAgent.name[0]}
                    </div>

                    {/* Информация об агенте */}
                    <div className="flex-1">
                        <h1 className="text-text-primary font-bold text-lg">
                            {currentAgent.name}
                        </h1>
                        <div className="flex items-center gap-2">
                            {/* Индикатор онлайн */}
                            <Circle className="w-2 h-2 fill-light-mint text-light-mint" />
                            <span className="text-text-secondary text-sm">
                                {currentAgent.status} • {currentAgent.personality}
                            </span>
                        </div>
                    </div>

                    {/* Дополнительные кнопки (меню, настройки) */}
                    <button className="
                        p-2
                        text-text-secondary
                        hover:text-text-primary
                        hover:bg-deep-midnight/50
                        rounded-lg
                        transition-colors
                    ">
                        <User className="w-5 h-5" />
                    </button>
                </header>

                {/* ========================================
                    MESSAGE LIST (История сообщений)
                    ======================================== */}
                <MessageList />

                {/* ========================================
                    MESSAGE INPUT (Поле ввода)
                    ======================================== */}
                <MessageInput />
            </div>
        </div>
    );
};