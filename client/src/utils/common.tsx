type VoidFunction = () => void;

type EventFunction = (event: React.ChangeEvent<HTMLInputElement>) => void;

export type { VoidFunction, EventFunction };

const onResponse = (res: any, func: () => void) => (!res.data.error ? func() : alert('Unexpected error :c'));

export { onResponse }; 
