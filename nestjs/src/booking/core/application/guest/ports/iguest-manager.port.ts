import { GuestRequest } from "../requests/guest.request";
import { GuestResponse } from "./../responses/guest.response";
export interface IGuestManager {
    Create(request: GuestRequest): Promise<GuestResponse>;
}