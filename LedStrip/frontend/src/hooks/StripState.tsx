import React, { useContext, useEffect, useState } from "react"

export type StripNumberProp = 'Brightness' | 'Speed'
export type StripProp = 'Color' | 'Mode' | 'Brightness' | 'Speed';

export type StripState = {
    Color: string
    Brightness: number
    Speed: number
    Mode: number
}

export type StripStateCtx = StripState & {
    setValue: (prop: StripProp, val: any) => void;
    useFavorite: (id: string) => void;
}

const initialState: StripState = {
    Color: 'FFFFFF',
    Brightness: 255,
    Speed: 1000,
    Mode: 0,
}

const StripStateContext = React.createContext<StripStateCtx>({
    ...initialState,

    setValue: (prop, val) => {},
    useFavorite: (id) => {},
})

export function StripStateProvider({children}: {children: React.ReactNode}) {
    const [state, setState] = useState<StripState>(initialState)

    useEffect(() => {
        const func = async () => {
            const resp = await fetch('/api/state')
            const json = await resp.json()

            setState(json)
        }

        func()
    }, [])

    const setValue = (prop: StripProp, val: any) => {
        setState({...state, [prop]: val})

        const fd = new FormData();
        fd.append(prop.toLowerCase(), val)

        fetch('/api/' + prop.toLowerCase() + '/set', {
            method: 'POST',
            body: fd,
        })
    }

    const useFavorite = (id: string) => {
        fetch('/api/favorite/' + id)
    }

    return <StripStateContext.Provider value={{
        ...state,
        setValue,
        useFavorite,
    }}>
        {children}
    </StripStateContext.Provider>
}

export function useLedStrip() {
    return useContext<StripStateCtx>(StripStateContext);
}