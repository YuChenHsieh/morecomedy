import React from 'react';
import style from './style'
import Box from '@mui/material/Box';
import { Swiper, SwiperSlide } from 'swiper/react';
import SwiperCore, {
    Navigation
} from 'swiper';

// install Swiper modules
SwiperCore.use([Navigation]);

// Import Swiper styles
import "swiper/css";
import "swiper/css/navigation"

const Event = (props) => {
    return (
        <Box sx={style.event}>
            <img src={props.src} alt="" style={style.image} />
        </Box>
    )
}

const Promo = (props) => {

    var Events = [
        <Event src='/comedybase.jpeg' />,
        <Event src='/burn.png' />,
        <Event src='/oldk.jpeg' />,
    ]

    return (
        <Swiper
            onSlideChange={() => console.log('slide change')}
            onSwiper={(swiper) => { }}
            navigation={true}
            style={style.carousel}
        >
            {Events.map(e => (<SwiperSlide>
                {e}
            </SwiperSlide>)
            )}
        </Swiper>
    );
}

export default Promo