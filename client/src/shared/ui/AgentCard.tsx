import type { AgentSummary } from '@/shared/types';

interface AgentCardProps {
    agent: AgentSummary;
    isSelected: boolean;
    onClick: (id: string) => void;
}

export const AgentCard = ({ agent, isSelected, onClick }: AgentCardProps) => {
    return (
        <div
            onClick={() => onClick(agent.id)}
            className={`
                group relative p-4 rounded-2xl cursor-pointer transition-all duration-300 border
                ${isSelected
                    ? 'bg-bright-turquoise/10 border-bright-turquoise/30 shadow-lg shadow-bright-turquoise/5'
                    : 'bg-white/[0.02] border-white/5 hover:border-white/10 hover:bg-white/[0.04]'
                }
            `}
        >
            <div className="flex items-center justify-between mb-2">
                <h3 className={`
                    text-sm font-bold transition-colors 
                    ${isSelected ? 'text-bright-turquoise' : 'text-white/80'}
                `}>
                    {agent.name}
                </h3>

                <div className="flex items-center gap-1.5">
                    <div className={`
                        w-1.5 h-1.5 rounded-full transition-all duration-500
                        ${agent.isActive 
                            ? 'bg-light-mint animate-pulse shadow-[0_0_8px_rgba(122,248,196,0.5)]' 
                            : 'bg-white/10'}
                    `} />
                    <span className="text-[10px] uppercase tracking-wider opacity-40 text-white font-medium">
                        {agent.isActive ? 'в сети' : 'спит'}
                    </span>
                </div>
            </div>

            <p className="text-white/40 text-xs line-clamp-1 font-light group-hover:text-white/60 transition-colors">
                {agent.personalityType}
            </p>

            <div className={`
                absolute inset-0 rounded-2xl transition-opacity duration-500 pointer-events-none
                ${isSelected ? 'opacity-100' : 'opacity-0 group-hover:opacity-100'}
                bg-gradient-to-br from-bright-turquoise/5 to-transparent
            `} />

            {isSelected && (
                <div className="absolute left-[-4px] top-1/4 bottom-1/4 w-1 bg-bright-turquoise rounded-full shadow-[0_0_15px_#26d0ce] animate-slide-up" />
            )}
        </div>
    );
};