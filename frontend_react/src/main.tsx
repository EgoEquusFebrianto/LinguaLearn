import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import App from './App.js'
import { RootProvider } from './RootProvider.js'

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <RootProvider>
      <App />
    </RootProvider>
  </StrictMode>,
)