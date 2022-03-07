import Card from 'components/common/Card'
import Box from '@mui/material/Box'
import style from './style'
import { useQuery } from '@apollo/client';
import graphql from '../../../graphql'

const Events = (props) => {
    const { loading, error, data } = useQuery(graphql.GET_EVENTS)
    if (loading) return (<Box></Box>)
    const events = data.events || []
    const eventsNum = events.length

    const shows = [
        {
            cover: '/logo.png',
            name: '週末夜喜劇秀',
            description: '「周末夜喜劇秀」Saturday Night Live Comedy\n大家一起來笑嗨嗨！'
        },
        {
            cover: '/logo.png',
            name: '週末夜喜劇秀',
            description: '「周末夜喜劇秀」Saturday Night Live Comedy\n大家一起來笑嗨嗨！'
        }, {
            cover: '/logo.png',
            name: '週末夜喜劇秀',
            description: '「周末夜喜劇秀」Saturday Night Live Comedy\n大家一起來笑嗨嗨！'
        },
        {
            cover: '/logo.png',
            name: '週末夜喜劇秀',
            description: '「周末夜喜劇秀」Saturday Night Live Comedy\n大家一起來笑嗨嗨！'
        },
        {
            cover: '/logo.png',
            name: '週末夜喜劇秀',
            description: '「周末夜喜劇秀」Saturday Night Live Comedy\n大家一起來笑嗨嗨！'
        }, {
            cover: '/logo.png',
            name: '週末夜喜劇秀',
            description: '「周末夜喜劇秀」Saturday Night Live Comedy\n大家一起來笑嗨嗨！'
        }
    ]
    const iteration = Array.from({ length: Math.ceil(eventsNum / 3) }, (v, i) => i)
    return (
        <>
            {iteration.map(i => (
                <>
                    <Box key={`event_row_${i}`} sx={style.row}>
                        {events.slice(i, i + 3).map((event) => (
                            <Box key={`event_${event.id}`} sx={style.col}>
                                <Card src={`/events/${event.id}.jpeg`} name={event.name} description={event.briefly} />
                            </Box>
                        ))}
                    </Box>
                </>
            ))
            }
        </>
    )
}

export default Events