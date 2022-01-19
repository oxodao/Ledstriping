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
}

const initialState: StripState = {
    Color: 'FFFFFF',
    Brightness: 255,
    Speed: 1000,
    Mode: 0,
}

const StripStateContext = React.createContext<StripStateCtx>({
    ...initialState,

    setValue: () => {},
})

export function StripStateProvider({children}: {children: React.ReactNode}) {
    const [state, setState] = useState<StripState>(initialState)

    // @TODO: useEffect to fetch initial state
    useEffect(() => {
        const func = async () => {
            const resp = await fetch('/api/state')
            const json = await resp.json()

            setState(json)
        }

        func()
    }, [])

    // @TODO: handle cors errors 
    const setValue = (prop: StripProp, val: any) => {
        setState({...state, [prop]: val})

        const fd = new FormData();
        fd.append(prop.toLowerCase(), val)

        fetch('/api/' + prop.toLowerCase() + '/set', {
            method: 'POST',
            body: fd,
        })
    }

    return <StripStateContext.Provider value={{
        ...state,
        setValue,
    }}>
        {children}
    </StripStateContext.Provider>
}

export function useLedStrip() {
    return useContext<StripStateCtx>(StripStateContext);
}