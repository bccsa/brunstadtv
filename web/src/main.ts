import { createApp } from "vue"
import "./styles/index.css"
import "./styles/barlow.css"
import "./styles/design-system.css"
import App from "./App.vue"
import router from "./router"
import auth0 from "@/services/auth0"

import { init as sentryInit, vueRouterInstrumentation } from "@sentry/vue"
import { BrowserTracing } from "@sentry/tracing"

import "swiper/css"
import "swiper/css/navigation"
import "swiper/css/pagination"
import "swiper/css/scrollbar"
import i18n from "./i18n"

const app = createApp(App)

sentryInit({
    app,
    dsn: "https://3ffd6244935a49dab6913bdc148d8d41@o1045703.ingest.sentry.io/4504299662278656",
    integrations: [
        new BrowserTracing({
            routingInstrumentation: vueRouterInstrumentation(router),
            tracePropagationTargets: [/^\//],
        }),
    ],
    tracesSampleRate: 1.0,
})

app.use(i18n).use(router).use(auth0)
app.mount("#app")
