import React from 'react';
import { render } from '@testing-library/react';
import App from './App';

test('renders hit button', () => {
  const { getByText } = render(<App />);
  const linkElement = getByText(/hit/i);
  expect(linkElement).toBeInTheDocument();
});
