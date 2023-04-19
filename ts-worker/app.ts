import * as grpc from '@grpc/grpc-js';
import { BooksService, BooksRequest, BooksResponse, BooksServer } from './books';

const host = '0.0.0.0:9090';
const server = new grpc.Server();

const booksServer: BooksServer = {
    get: (call, callback) => {
        console.log(`(server) Got client message: ${call.request.query}`);

        // logic here
        const response: BooksResponse = {
            books: []
        };

        callback(null, response);
    }
};
server.addService(BooksService, booksServer);

server.bindAsync(
    host,
    grpc.ServerCredentials.createInsecure(),
    (err: Error | null, port: number) => {
        if (err) {
            console.error(`Server error: ${err.message}`);
        } else {
            console.log(`Server bound on port: ${port}`);
            server.start();
        }
    }
);
