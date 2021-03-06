import React, { useContext, useEffect, useState } from "react"
import { Favorite } from "./DataState";

export type StripNumberProp = 'Brightness' | 'Speed'
export type StripProp = 'Color' | 'Mode' | 'Brightness' | 'Speed';

export type StripState = {
    Color: string
    Brightness: number
    Speed: number
    Mode: string
}

export type StripStateCtx = StripState & {
    setValue: (prop: StripProp, val: any) => void;
    useFavorite: (id: Favorite) => void;
}

const initialState: StripState = {
    Color: 'FFFFFF',
    Brightness: 255,
    Speed: 1000,
    Mode: "Static",
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

    const useFavorite = (favorite: Favorite) => {
        fetch('/api/favorite/' + favorite.ID)
        setState({
            ...state,
            Color: favorite.Color,
            Brightness: favorite.Brightness,
            Mode: favorite.Mode,
            Speed: favorite.Speed,
        })
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