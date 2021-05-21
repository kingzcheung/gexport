import axios from "axios";

axios.defaults.baseURL = `${import.meta.env.VITE_APP_API}/api`;
// axios.defaults.headers.common['Authorization'] = "Bearer " + token;
const _axios = axios.create({});
_axios.interceptors.request.use(
    function (config) {
        // Do something before request is sent
        return config;
    },
    function (error) {
        // Do something with request error
        return Promise.reject(error);
    }
);

// Add a response interceptor
_axios.interceptors.response.use(
    function (response) {
        // Do something with response data
        return response;
    },
    function (error) {
        // Do something with response error
        if (null !== error) {
            console.log(error.response)
            if (error.response.status === 401) {
                window.location.href = `${import.meta.env.VITE_APP_API}/auth/oauth`
                return
            }
        }
        return Promise.reject(error);
    }
);


export default _axios;