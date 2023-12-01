import { PUBLIC_API_ENDPOINT } from "$env/static/public";

const endpoint = () => `${PUBLIC_API_ENDPOINT}/counter`

type ResponseCount = {
    count: number;
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