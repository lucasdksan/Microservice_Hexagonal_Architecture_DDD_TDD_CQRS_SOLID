export default class RoomEntity {
    id: number;
    name: string;
    level: number;
    inMaintenance: boolean;

    isAvailable(): boolean {
        return !this.inMaintenance || this.hasGuest() ? false : true;
    }

    hasGuest(): boolean {
        return true;
    }
}