import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Image from 'next/image'
import style from './style'
import SearchButton from '../../common/SearchButton';

export default function HeaderAppBar() {
  return (
        <AppBar sx={style.appbar} position="static">
         <Toolbar sx={style.toolbar}>
          <Box sx={style.logo}>
            <Image src='/logo_words.jpg' width={108} height={48}/>
          </Box>
          <SearchButton/>
         </Toolbar>
        </AppBar>
  );
}
