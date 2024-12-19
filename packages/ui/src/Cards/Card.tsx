import { cn } from "@oauth-service-go/utils";

interface CardProps {
  children?: React.ReactNode;
  className?: string;
}

export function Card({ children, className }: CardProps) {
  return (
    <div
      className={cn(
        "rounded-xl border-2",
        "shadow-lg shadow-slate-300",
        "m-4 p-4",
        className,
      )}
    >
      {children}
    </div>
  );
}
