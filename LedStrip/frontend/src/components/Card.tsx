import React from "react";

import "../assets/css/card.scss";

type CardProps = {
    title: string
    children: React.ReactNode
    className?: string
}

export default function Card({title, children, className}: CardProps) {
    return <div className={`Card${className ? ' ' + className : ''}`}>
        <h1>{title}</h1>
        <div className="Card__Content">
            {children}
        </div>
    </div>
}