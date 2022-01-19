import React, { useContext, useEffect, useState } from "react"

const APP_NAME = 'ledstrip-web-ui'

export type Favorite = {
    ID?: string
    App?: string
    Name: string
    Color: string
    Brightness: number
    Mode: number
    Speed: number
}

export type DataState = {
    Modes: Array<string>,
    Favorites: Array<Favorite>

    currentName: string
}

export type DataStateCtx = DataState & {
    refresh: () => void;
    addFavorite: (f: Favorite) => void;
    setCurrentName: (s: string) => void;
}

const initialState: DataState = {
    Modes: [],
    Favorites: [],

    currentName: '',
}

const DataStateContext = React.createContext<DataStateCtx>({
    ...initialState,
    refresh: () => { },
    addFavorite: () => { },
    setCurrentName: () => { },
})

export function DataStateProvider({ children }: { children: React.ReactNode }) {
    const [state, setState] = useState<DataState>(initialState)

    const refresh = () => {
        const func = async () => {
            const resp = await fetch('/api/data')
            const json = await resp.json()

            json['Favorites'] = json['Favorites'].filter((i: Favorite) => i.App === APP_NAME)

            setState(json)
        }

        func()
    }

    const setCurrentName = (currentName: string) => {
        setState({ ...state, currentName })
    }

    const addFavorite = async (obj: Favorite) => {
        if (!obj.Name || obj.Name.length === 0) {
            return
        }

        obj.App = APP_NAME;

        await fetch('/api/favorite', {
            method: 'POST',
            body: JSON.stringify(obj)
        })
        await refresh()

        setState({ ...state, currentName: '' })
    }

    useEffect(() => {
        refresh()
    }, [])


    return <DataStateContext.Provider value={{
        ...state,
        refresh,
        addFavorite,
        setCurrentName,
    }}>
        {children}
    </DataStateContext.Provider>
}

export function useData() {
    return useContext<DataStateCtx>(DataStateContext);
}