/**
 * MessageInput - Поле ввода сообщения
 * 
 * Этот компонент отображает:
 * - Текстовое поле для ввода сообщения
 * - Кнопку отправки
 * - Индикатор количества символов (опционально)
 * 
 * TODO для декомпозиции:
 * - Добавить поддержку multiline (textarea)
 * - Добавить кнопки для прикрепления файлов
 * - Добавить эмодзи-пикер
 * - Добавить автозаполнение
 */

import { Send } from 'lucide-react';

export const MessageInput = () => {
    return (
        <div className="
            /* Контейнер для поля ввода */
            border-t border-bright-turquoise/20
            bg-dark-ocean
            p-4
            /* Sticky footer на мобилке */
            sticky bottom-0
        ">
            {/* ФОРМА ВВОДА */}
            <form className="
                flex items-end gap-3
                max-w-4xl mx-auto
            ">
                {/* ТЕКСТОВОЕ ПОЛЕ */}
                <div className="flex-1">
                    <textarea
                        placeholder="Type your message to Alice..."
                        rows={1}
                        className="
                            /* Базовые стили */
                            w-full
                            px-4 py-3
                            bg-deep-midnight
                            border border-bright-turquoise/30
                            rounded-xl
                            text-text-primary
                            placeholder:text-text-secondary/50
                            
                            /* Фокус */
                            focus:outline-none
                            focus:ring-2
                            focus:ring-bright-turquoise
                            focus:border-transparent
                            
                            /* Адаптивность */
                            text-sm md:text-base
                            
                            /* Автоматическое изменение высоты */
                            resize-none
                            min-h-[44px]
                            max-h-[120px]
                            
                            /* Скроллбар */
                            scrollbar-thin scrollbar-thumb-bright-turquoise/30 scrollbar-track-transparent
                        "
                    />
                </div>

                {/* КНОПКА ОТПРАВКИ */}
                <button
                    type="submit"
                    className="
                        /* Базовые стили */
                        px-4 py-3
                        bg-gradient-primary
                        text-white
                        rounded-xl
                        
                        /* Hover эффект */
                        hover:shadow-lg
                        
                        /* Фокус */
                        focus:outline-none
                        focus:ring-2
                        focus:ring-bright-turquoise
                        focus:ring-offset-2
                        focus:ring-offset-deep-midnight
                        
                        /* Transition */
                        transition-all
                        
                        /* Disabled состояние (когда поле пустое) */
                        disabled:opacity-50
                        disabled:cursor-not-allowed
                        
                        /* Адаптивность */
                        flex items-center justify-center
                        min-w-[44px]
                    "
                >
                    {/* Иконка отправки */}
                    <Send className="w-5 h-5" />

                    {/* Текст кнопки (скрывается на мобилке) */}
                    <span className="ml-2 hidden md:inline font-semibold">
                        Send
                    </span>
                </button>
            </form>

            {/* ДОПОЛНИТЕЛЬНАЯ ИНФОРМАЦИЯ (опционально) */}
            {/* 
            <div className="mt-2 flex justify-between items-center text-xs text-text-secondary/70">
                <span>Press Enter to send, Shift+Enter for new line</span>
                <span>0 / 500</span>
            </div>
            */}
        </div>
    );
};
