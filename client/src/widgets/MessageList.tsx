/**
 * MessageList - Список сообщений в чате
 * 
 * Этот компонент отображает:
 * - Историю сообщений между пользователем и AI-агентом
 * - Разделение на сообщения от пользователя и от AI
 * - Автоскролл к последнему сообщению
 * 
 * TODO для декомпозиции:
 * - Вынести Message в отдельный компонент (UserMessage, AIMessage)
 * - Добавить индикатор "печатает..."
 * - Добавить временные метки
 * - Добавить аватары
 */

export const MessageList = () => {
    // Моковые данные сообщений (потом заменить на API)
    const messages = [
        {
            id: '1',
            sender: 'ai',
            senderName: 'Alice',
            content: 'Hello! I\'m Alice, a curious AI agent. How can I help you today?',
            timestamp: '10:30 AM',
        },
        {
            id: '2',
            sender: 'user',
            senderName: 'You',
            content: 'Hi Alice! Can you tell me about your personality?',
            timestamp: '10:31 AM',
        },
        {
            id: '3',
            sender: 'ai',
            senderName: 'Alice',
            content: 'I\'m characterized by high openness and curiosity. I love exploring new ideas and asking questions. My core values include honesty and creativity.',
            timestamp: '10:31 AM',
        },
        {
            id: '4',
            sender: 'user',
            senderName: 'You',
            content: 'That\'s interesting! What are you thinking about right now?',
            timestamp: '10:32 AM',
        },
    ];

    return (
        <div className="
            /* Контейнер для сообщений */
            flex-1
            overflow-y-auto
            p-4 md:p-6
            space-y-4
            /* Кастомный скроллбар */
            scrollbar-thin scrollbar-thumb-bright-turquoise/30 scrollbar-track-transparent
        ">
            {messages.map((message) => (
                <div
                    key={message.id}
                    className={`
                        /* Выравнивание: AI слева, пользователь справа */
                        flex
                        ${message.sender === 'user' ? 'justify-end' : 'justify-start'}
                    `}
                >
                    {/* КОНТЕЙНЕР СООБЩЕНИЯ */}
                    <div className={`
                        /* Максимальная ширина сообщения */
                        max-w-[85%] md:max-w-[70%]
                        flex flex-col
                        ${message.sender === 'user' ? 'items-end' : 'items-start'}
                    `}>
                        {/* ИМЯ ОТПРАВИТЕЛЯ */}
                        <div className="
                            flex items-center gap-2 mb-1
                            px-2
                        ">
                            {/* Индикатор онлайн (только для AI) */}
                            {message.sender === 'ai' && (
                                <div className="w-2 h-2 bg-light-mint rounded-full" />
                            )}

                            <span className="text-text-secondary text-xs font-medium">
                                {message.senderName}
                            </span>

                            <span className="text-text-secondary/50 text-xs">
                                {message.timestamp}
                            </span>
                        </div>

                        {/* ТЕЛО СООБЩЕНИЯ */}
                        <div className={`
                            px-4 py-3
                            rounded-2xl
                            shadow-md
                            
                            /* Разные стили для AI и пользователя */
                            ${message.sender === 'ai'
                                ? 'bg-dark-ocean text-text-primary rounded-tl-none'
                                : 'bg-gradient-primary text-white rounded-tr-none'
                            }
                        `}>
                            <p className="text-sm md:text-base leading-relaxed">
                                {message.content}
                            </p>
                        </div>
                    </div>
                </div>
            ))}

            {/* ИНДИКАТОР "ПЕЧАТАЕТ..." (опционально, показывать когда AI генерирует ответ) */}
            {/* 
            <div className="flex justify-start">
                <div className="bg-dark-ocean px-4 py-3 rounded-2xl rounded-tl-none">
                    <div className="flex gap-1">
                        <div className="w-2 h-2 bg-bright-turquoise rounded-full animate-bounce" style={{ animationDelay: '0ms' }} />
                        <div className="w-2 h-2 bg-bright-turquoise rounded-full animate-bounce" style={{ animationDelay: '150ms' }} />
                        <div className="w-2 h-2 bg-bright-turquoise rounded-full animate-bounce" style={{ animationDelay: '300ms' }} />
                    </div>
                </div>
            </div>
            */}
        </div>
    );
};
