// import {PUBLIC_API_ENDPOINT} from "$env/static/public";

// const endpoint = () => `${PUBLIC_API_ENDPOINT}/counter`
const endpoint = () => "https://b4qocpsyq5.execute-api.ap-northeast-1.amazonaws.com/counter"
type ResponseCount = {
    Count: number;
};

export const getCount = async (): Promise<ResponseCount> => {
    const res = await fetch(endpoint(), {
        method: 'GET',
        mode: 'cors',
    });
    return res.json();
};

export const updateCount = async (): Promise<ResponseCount> => {
    const res = await fetch(endpoint(), {
        method: 'POST',
        mode: 'cors',
    });
    return res.json();
};