export namespace AppResponse {
    export enum ErrorCodes {
        NOT_FOUND = 1,
        COULDNOT_STORE_DATA = 2,
    }

    export abstract class Response {
        success: boolean;
        message: string;
        errorCode?: ErrorCodes;
    }
}