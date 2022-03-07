import { ApolloClient, InMemoryCache } from '@apollo/client';
import { GET_EVENTS } from './events'
import config from 'config'

const client = new ApolloClient({
    uri: config.GRAPHQL_URL,
    cache: new InMemoryCache()
});

export default { client, GET_EVENTS }