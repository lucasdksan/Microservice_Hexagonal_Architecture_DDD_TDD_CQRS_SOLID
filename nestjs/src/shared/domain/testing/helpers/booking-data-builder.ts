import { faker } from "@faker-js/faker";
import { BookingProps } from "@/booking/core/domain/entities/booking.entity";

type Props = {
    id?: number;
    start?: Date;
    end?: Date;
    placedAt?: Date;
}

function generateRandomId(min: number = 1, max: number = 10000): number {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

export function BookingDataBuilder(props: Props): BookingProps {
    return {
        id: props.id ??  generateRandomId(),
        start: props.start ?? faker.date.future(),
        end: props.end ?? faker.date.future(),
        placedAt: props.placedAt ?? new Date(),
    };
}
