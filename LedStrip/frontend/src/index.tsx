import React from 'react';
import ReactDOM from 'react-dom';
import ColorSelector from './components/ColorSelector';
import FavoriteBrowser from './components/FavoriteBrowser';
import ModeSelector from './components/ModeSelector';
import { BrightnessSelector, SpeedSelector } from './components/Sliders';
import { DataStateProvider } from './hooks/DataState';
import { StripStateProvider } from './hooks/StripState';

import './assets/css/index.scss';
import 'font-awesome/css/font-awesome.min.css';

ReactDOM.render(
  <React.StrictMode>
    <DataStateProvider>
      <StripStateProvider>
        <FavoriteBrowser />
        <ColorSelector />
        <BrightnessSelector />
        <SpeedSelector />
        <ModeSelector />
      </StripStateProvider>
    </DataStateProvider>
  </React.StrictMode>,
  document.getElementById('root')
);
