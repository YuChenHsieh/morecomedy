import Head from 'next/head'
import styles from '../styles/Home.module.css'
import Header from 'components/Home/Header'
import Promo from 'components/Home/Promo'
import Events from 'components/Home/Events'
import Box from '@mui/material/Box'

export default function Home () {
  return (
    <Box sx={{ overflow: 'auto' }}>
      <Head>
        <title>Comedy TW</title>
        <meta name='description' content='Comedy Shows in Taiwan' />
        <link rel='icon' href='/microphone.png' />
      </Head>

      <main className={styles.main}>
        <Box sx={{
          width: '100%', height: '10%',
          display: 'flex',
          flexDirection: 'column',
        }}>
          <Header />
        </Box>
        <Box sx={{ display: 'flex', flexDirection: 'column', width: '100%', height: '50vh' }}>
          <Promo />
        </Box>
        <Box sx={{
          display: 'flex', flexDirection: 'column', width: '80%', marginLeft: 'auto', marginRight: 'auto',
        }}>
          <Events />
        </Box>
      </main>

      <footer className={styles.footer}>
        Powered by ComedyTW
      </footer>
    </Box>
  )
}
