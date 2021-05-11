import React              from 'react';
import ReactDOM           from 'react-dom';
import App                                from './App';
import {createMuiTheme, MuiThemeProvider} from "@material-ui/core";

ReactDOM.render(
  <React.StrictMode>
      <MuiThemeProvider theme={createMuiTheme({
          palette: {
              type: 'dark',
              primary: {
                main: '#2d7c9d'
              },
          },
      })}>
        <App />
      </MuiThemeProvider>
  </React.StrictMode>,
  document.getElementById('root')
);
