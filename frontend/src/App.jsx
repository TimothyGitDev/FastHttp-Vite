import { useEffect, useState } from 'react'

function App() {
    const [data, setData] = useState('');

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch('http://localhost:1234/api');
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                const result = await response.text();
                setData(result);
            } catch (error) {
                console.error('Ошибка при получении данных:', error);
            }
        };

        fetchData();
    }, []);

    return (
        <div>
            <h1>Данные с бэкэнда:</h1>
            <p>{data}</p>
        </div>
    );
}

export default App;