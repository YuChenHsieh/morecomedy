import { gql } from '@apollo/client';

const GET_EVENTS = gql`
query{
    events{
      id
      name
      briefly
      category
      host{
        name
      }
      shows{
        price
        location
      }
    }
  }
`

module.exports = { GET_EVENTS }