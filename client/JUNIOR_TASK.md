# Задание для Junior разработчика: Создание компонента AgentCard

## 📋 Описание задачи

Создать переиспользуемый компонент **AgentCard** — карточку AI-агента для отображения в списке.

## 🎯 Цель

Научиться создавать изолированные UI-компоненты в Storybook с использованием кастомной цветовой палитры проекта.

## 📍 Где создавать

```
client/src/shared/ui/AgentCard/
├── AgentCard.tsx
└── AgentCard.stories.tsx
```

## 🎨 Дизайн компонента

Карточка должна отображать:
- **Аватар** — первая буква имени агента на градиентном фоне
- **Имя агента** — крупным шрифтом
- **Тип личности** — мелким текстом
- **Индикатор статуса** — зеленая точка (онлайн) или серая (оффлайн)
- **Hover эффект** — тень при наведении

### Макет

```
┌─────────────────────────────────┐
│  ┌───┐  Alice          ● online │
│  │ A │  Curious Explorer        │
│  └───┘                           │
└─────────────────────────────────┘
```

## 📝 Интерфейс компонента

```tsx
interface AgentCardProps {
  name: string;
  personality: string;
  status: 'online' | 'offline';
  onClick?: () => void;
}
```

## 🎨 Цвета (из палитры проекта)

- Фон карточки: `bg-dark-ocean`
- Текст имени: `text-text-primary`
- Текст личности: `text-text-secondary`
- Аватар: `bg-gradient-primary`
- Индикатор онлайн: `bg-light-mint`
- Индикатор оффлайн: `bg-text-secondary/50`
- Hover: `hover:shadow-lg`

## 📂 Шаг 1: Создать файл компонента

**Путь:** `client/src/shared/ui/AgentCard/AgentCard.tsx`

```tsx
interface AgentCardProps {
  name: string;
  personality: string;
  status: 'online' | 'offline';
  onClick?: () => void;
}

export const AgentCard = ({ name, personality, status, onClick }: AgentCardProps) => {
  return (
    <div
      onClick={onClick}
      className="
        bg-dark-ocean
        rounded-xl
        p-4
        cursor-pointer
        transition-all
        hover:shadow-lg
        flex items-center gap-4
      "
    >
      {/* Аватар */}
      <div className="
        w-12 h-12
        bg-gradient-primary
        rounded-full
        flex items-center justify-center
        text-white text-xl font-bold
      ">
        {name[0]}
      </div>

      {/* Информация */}
      <div className="flex-1">
        <h3 className="text-text-primary font-semibold text-lg">
          {name}
        </h3>
        <p className="text-text-secondary text-sm">
          {personality}
        </p>
      </div>

      {/* Статус */}
      <div className="flex items-center gap-2">
        <div className={`
          w-3 h-3 rounded-full
          ${status === 'online' ? 'bg-light-mint animate-pulse' : 'bg-text-secondary/50'}
        `} />
        <span className="text-text-secondary text-xs">
          {status}
        </span>
      </div>
    </div>
  );
};
```

## 📂 Шаг 2: Создать Storybook stories

**Путь:** `client/src/shared/ui/AgentCard/AgentCard.stories.tsx`

```tsx
import type { Meta, StoryObj } from '@storybook/react';
import { AgentCard } from './AgentCard';

const meta: Meta<typeof AgentCard> = {
  title: 'Shared/AgentCard',
  component: AgentCard,
  tags: ['autodocs'],
  argTypes: {
    status: {
      control: 'select',
      options: ['online', 'offline'],
    },
  },
};

export default meta;
type Story = StoryObj<typeof AgentCard>;

export const Online: Story = {
  args: {
    name: 'Alice',
    personality: 'Curious Explorer',
    status: 'online',
  },
};

export const Offline: Story = {
  args: {
    name: 'Bob',
    personality: 'Wise Guardian',
    status: 'offline',
  },
};

export const Interactive: Story = {
  args: {
    name: 'Charlie',
    personality: 'Creative Dreamer',
    status: 'online',
    onClick: () => alert('Agent clicked!'),
  },
};
```

## 📂 Шаг 3: Добавить экспорт

**Путь:** `client/src/shared/ui/AgentCard/index.ts`

```tsx
export { AgentCard } from './AgentCard';
```

## 🚀 Команды для запуска

### 1. Запустить Storybook

```bash
cd /home/tima/Desktop/milk-island/client
npm run storybook
```

### 2. Открыть в браузере

```
http://localhost:6006
```

### 3. Найти компонент

В левой панели Storybook:
```
Shared → AgentCard
```

## ✅ Критерии приемки

- [ ] Компонент создан в `client/src/shared/ui/AgentCard/AgentCard.tsx`
- [ ] Stories созданы в `client/src/shared/ui/AgentCard/AgentCard.stories.tsx`
- [ ] Экспорт добавлен в `client/src/shared/ui/AgentCard/index.ts`
- [ ] Компонент отображается в Storybook
- [ ] Все 3 истории работают (Online, Offline, Interactive)
- [ ] Используются цвета из палитры проекта
- [ ] Hover эффект работает
- [ ] Индикатор онлайн анимируется (pulse)

## 📚 Полезные ссылки

- [Документация цветовой палитры](file:///home/tima/Desktop/milk-island/CONTRIBUTING.md#цветовая-палитра-проекта)
- [Tailwind CSS документация](https://tailwindcss.com/docs)
- [Storybook документация](https://storybook.js.org/docs)

## 💡 Подсказки

1. **Первая буква имени:** `name[0]` или `name.charAt(0)`
2. **Условные классы:** используй тернарный оператор `${condition ? 'class-a' : 'class-b'}`
3. **Анимация pulse:** уже есть в Tailwind, просто добавь `animate-pulse`
4. **Проверка в Storybook:** используй Controls для изменения пропсов в реальном времени

## 🎓 Что ты изучишь

- ✅ Создание переиспользуемых компонентов
- ✅ Работа с TypeScript интерфейсами
- ✅ Использование Tailwind CSS
- ✅ Работа со Storybook
- ✅ Условный рендеринг в React

## ⏱️ Примерное время выполнения

**30-45 минут**

---

**Удачи! Если возникнут вопросы — обращайся! 🚀**
