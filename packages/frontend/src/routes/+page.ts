import type { PageLoad } from './$types';
import { getCount } from "$lib";

export const load: PageLoad = async () => {
    return await getCount();
};