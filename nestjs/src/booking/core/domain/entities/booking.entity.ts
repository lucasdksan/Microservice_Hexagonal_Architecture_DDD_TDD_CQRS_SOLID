import { Actions } from "../enums/actions.enum";
import { Status } from "../enums/status.enum";

export default class BookingEntity {
    public id: number;
    public placedAt: Date;
    public start: Date;
    public end: Date;
    private status: Status

    constructor() {
        this.status = Status.Created;
    }

    currentStatus(): Status {
        return this.status;
    }

    changeState(action: Actions): void {
        switch (this.status) {
            case Status.Created:
                if (action === Actions.Pay) {
                    this.status = Status.Paid;
                } else if (action === Actions.Cancel) {
                    this.status = Status.Canceled;
                }
                break;

            case Status.Paid:
                if (action === Actions.Finish) {
                    this.status = Status.Finished;
                } else if (action === Actions.Refound) {
                    this.status = Status.Refounded;
                }
                break;

            case Status.Canceled:
                if (action === Actions.Reopen) {
                    this.status = Status.Created;
                }
                break;
        }
    }
}