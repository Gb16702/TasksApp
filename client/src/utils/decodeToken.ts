import jwt from 'jsonwebtoken';

const decodeToken = async (token: string) => {
    const value = await jwt.decode(token);

    return value;
}

export default decodeToken;
