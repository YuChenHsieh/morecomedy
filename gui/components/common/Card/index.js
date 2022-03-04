import Card from '@mui/material/Card';
import CardActionArea from '@mui/material/CardActionArea';
import CardMedia from '@mui/material/CardMedia';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';

const EventCard = (props) => {
    const { src, name, description } = props
    return (
        <Card sx={{ width: '100%', height: '100%', borderRadius: '10px' }}>
            <CardActionArea>
                <CardMedia
                    component="img"
                    height="200px"
                    image={src}
                    alt="event"
                />
                <CardContent>
                    <Typography gutterBottom variant="h5" component="div">
                        {name}
                    </Typography>
                    <Typography variant="body2" color="text.secondary">
                        {description}
                    </Typography>
                </CardContent>
            </CardActionArea>
        </Card>
    );
}

export default EventCard
