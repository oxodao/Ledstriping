import React from 'react';
import ReactDOM from 'react-dom';

import './assets/css/index.scss';
import ColorSelector from './components/ColorSelector';
import ModeSelector from './components/ModeSelector';
import { BrightnessSelector, SpeedSelector } from './components/Sliders';
import { StripStateProvider } from './hooks/StripState';

ReactDOM.render(
  <React.StrictMode>
    <StripStateProvider>
      <ColorSelector />
      <BrightnessSelector />
      <SpeedSelector />
      <ModeSelector />
    </StripStateProvider>
  </React.StrictMode>,
  document.getElementById('root')
);
