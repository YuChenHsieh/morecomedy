import '../styles/globals.css'
import {
  ApolloProvider,
} from "@apollo/client";
import graphql from '../graphql'

function MyApp ({ Component, pageProps }) {
  return (
    <ApolloProvider client={graphql.client}>
      <Component {...pageProps} />
    </ApolloProvider>
  )
}

export default MyApp
