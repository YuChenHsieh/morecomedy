import Card from 'components/common/Card'
import Box from '@mui/material/Box'
import style from './style'
import { useQuery } from '@apollo/client';
import { GET_EVENTS } from 'pages/graphql'

const Events = (props) => {
    const { loading, error, data } = useQuery(GET_EVENTS)
    console.log('data', data)
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

    return (
        <>
            <Box sx={style.row}>
                {shows.slice(0, 3).map((show, index) => (
                    <Box key={`show_${index}`} sx={style.col}>
                        <Card src={show.cover} name={show.name} description={show.description} />
                    </Box>
                ))}
            </Box>
            <Box sx={style.row}>
                {shows.slice(3, 6).map((show, index) => (
                    <Box key={`show_${index + 3}`} sx={style.col}>
                        <Card src={show.cover} name={show.name} description={show.description} />
                    </Box>
                ))}
            </Box>
        </>
    )
}

export default Events