# Contributing to MindFlow (Frontend)

Welcome! This guide will help you get started with frontend development for the project.

## Project Architecture

We utilize a **Feature-Based Architecture** (inspired by Feature-Sliced Design). The code is organized by business domains rather than technical types.

### Directory Structure

```
src/
├── app/          # Global application configuration
├── pages/        # Routing pages (wrappers)
├── features/     # Business logic and specific UI
├── shared/       # Universal reusable components
├── widgets/      # Large composite UI blocks
└── entities/     # Business entities (data models)

```

#### Detailed Layer Description

**`app/`** — Global application configuration

* `app/router/` — Routing configuration (React Router)
* `app/providers/` — Global providers (QueryClient, AuthProvider)
* `app/styles/` — Global styles

**`pages/`** — Page components for routing

* Contains **minimal logic**, focusing solely on component composition
* Example: `pages/auth/ui/LoginPage.tsx` — simply places the `LoginForm` on the page
* Structure: `pages/<domain>/ui/Page.tsx`

**`features/`** — Business logic and specific UI components

* Each feature is isolated and contains everything necessary for its operation
* Feature structure:
```
features/auth/
├── ui/               # UI components specific to authentication
│   └── LoginForm/    # Login form (aware of email/password fields)
├── model/            # Business logic (hooks, state)
└── api/              # API requests for this specific feature

```


* **Important**: UI in `features` consists of components specific to that particular feature.

**`shared/`** — Universal reusable components

* Contains code that **DOES NOT** know about business logic
* Structure:
```
shared/
├── ui/        # Universal UI components (Button, Input, Card)
├── api/       # Base API client (axios instance)
├── hooks/     # Universal hooks (useDebounce, useLocalStorage)
└── lib/       # Utilities (cn, formatters, validators)

```


* **Important**: Components in `shared/ui` must be as universal as possible.

**`widgets/`** — Large composite UI blocks

* Complex components consisting of several `shared` or `feature` components
* Examples: `Navbar`, `Sidebar`, `Footer`
* Can use components from `shared`, but should avoid direct dependencies on unrelated `features` when possible.

**`entities/`** — Business Entities

* Data models (Agent, User, Relationship)
* TypeScript types, validation schemas

#### Why is `auth` present in both `features` and `pages`?

**They serve different purposes:**

1. **`features/auth/`** — Contains the **authentication business logic**:
* `features/auth/ui/LoginForm/` — The login form (specific UI)
* `features/auth/model/` — Hooks for handling authentication
* `features/auth/api/` — Authentication API requests


2. **`pages/auth/`** — Contains **wrapper pages** for routing:
* `pages/auth/ui/LoginPage.tsx` — Simply assembles the `LoginForm` on the page
* Minimal logic, focusing on composition only



**Analogy:** `features` are the "bricks" containing logic; `pages` are the "walls" assembled from those bricks.

#### Why is there a `ui` folder in `features` if `shared/ui` exists?

**Difference in intent:**

* **`shared/ui/Button`** — A universal button, unaware of the business context.
```tsx
// Can be used anywhere
<Button onClick={...}>Click Me</Button>

```


* **`features/auth/ui/LoginForm`** — A login form, aware of email/password logic.
```tsx
// Specific to authentication, uses shared components internally
<LoginForm onSubmit={handleLogin} />

```



**The Rule:** If a component is specific to one feature, it belongs in `features/*/ui`. If it is universal, it belongs in `shared/ui`.

### Key Technologies

* **Vite**: Build tool.
* **React**: Library for building user interfaces.
* **Tailwind CSS**: Styling.
* **Storybook**: Component development in an isolated environment.
* **TypeScript**: Ensuring type safety.
* **Lucide React**: Project iconography.

## Project Color Palette

The project uses a custom color palette configured in `tailwind.config.ts`. All colors are accessible via Tailwind CSS classes.

### 🔵 Deep Blues (Background and Primary Blocks)

These colors evoke a sense of stability and professionalism.

| Color | HEX | Tailwind Class | Purpose |
| --- | --- | --- | --- |
| Deep Midnight Blue | `#0B1E3B` | `bg-deep-midnight` | Main application background |
| Dark Ocean Blue | `#1A3C5E` | `bg-dark-ocean` | Cards, secondary elements |

**Usage Examples:**

```tsx
// Main app background (already applied to body)
<div className="bg-deep-midnight">...</div>

// Card with dark background
<div className="bg-dark-ocean rounded-lg p-6">
  <h2 className="text-text-primary">Card Title</h2>
</div>

```

### 💎 Blues and Cyans (Accents and Active Elements)

Colors derived from the top of the logo gradient, ideal for buttons and links.

| Color | HEX | Tailwind Class | Purpose |
| --- | --- | --- | --- |
| Bright Turquoise | `#26D0CE` | `bg-bright-turquoise` | Main accent color |
| Sky Blue | `#5BC0EB` | `bg-sky-blue` | Icons, secondary buttons |

**Usage Examples:**

```tsx
// Primary button with accent color
<button className="bg-bright-turquoise hover:bg-sky-blue text-white px-6 py-3 rounded-lg transition-colors">
  Login
</button>

// Link with accent color
<a href="#" className="text-bright-turquoise hover:text-sky-blue">
  Learn More
</a>

// Icon with sky blue color
<Icon className="text-sky-blue" />

```

### 🟢 Greens and Mints (Processes and AI "Thoughts")

Symbolize system health and successful task execution.

| Color | HEX | Tailwind Class | Purpose |
| --- | --- | --- | --- |
| Light Mint | `#7AF8C4` | `bg-light-mint` | "Online" indicators, success |
| Soft Teal | `#50E3C2` | `bg-soft-teal` | Chat separation, text highlighting |

**Usage Examples:**

```tsx
// "Online" status indicator
<div className="flex items-center gap-2">
  <div className="w-3 h-3 bg-light-mint rounded-full"></div>
  <span className="text-text-secondary">Online</span>
</div>

// Success notification
<div className="bg-soft-teal/20 border border-soft-teal rounded-lg p-4">
  <p className="text-soft-teal">Agent created successfully!</p>
</div>

// Active chat highlight
<div className="border-l-4 border-soft-teal bg-dark-ocean p-4">
  <p>Active Chat</p>
</div>

```

### 📝 Text Colors

| Color | HEX | Tailwind Class | Purpose |
| --- | --- | --- | --- |
| White | `#FFFFFF` | `text-text-primary` | Headings, primary text |
| Light Gray | `#B0BEC5` | `text-text-secondary` | Secondary text, descriptions |

**Usage Examples:**

```tsx
// Heading
<h1 className="text-text-primary text-3xl font-bold">Heading</h1>

// Description
<p className="text-text-secondary text-sm">Additional information</p>

```

### 🎨 Gradients

The project includes two ready-to-use gradients:

| Name | Tailwind Class | Purpose |
| --- | --- | --- |
| Primary Gradient | `bg-gradient-primary` | Buttons, accents |
| Accent Gradient | `bg-gradient-accent` | Special elements |

**Usage Examples:**

```tsx
// Button with gradient (matches favicon style)
<button className="bg-gradient-primary text-white px-8 py-4 rounded-lg font-semibold shadow-lg hover:shadow-xl transition-shadow">
  Create Agent
</button>

// Card with gradient background
<div className="bg-gradient-accent p-6 rounded-xl">
  <h3 className="text-deep-midnight font-bold">Special Offer</h3>
</div>

// Gradient text
<h1 className="bg-gradient-primary bg-clip-text text-transparent text-5xl font-bold">
  Milk Island AI
</h1>

```

### 🎯 Usage Recommendations

1. **App Background**: Always use `bg-deep-midnight` (already applied to `body`).
2. **Cards and Blocks**: `bg-dark-ocean` for content separation.
3. **Buttons**: `bg-gradient-primary` for primary actions, `bg-bright-turquoise` for secondary.
4. **Text**: `text-text-primary` for headings, `text-text-secondary` for descriptions.
5. **Statuses**: `bg-light-mint` for success/online, `bg-soft-teal` for active elements.
6. **Links**: `text-bright-turquoise hover:text-sky-blue`.
